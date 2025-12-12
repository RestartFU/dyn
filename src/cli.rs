use crate::logger;
use owo_colors::OwoColorize;
use std::env;
use std::fs;
use std::io::{BufRead, BufReader};
use std::os::unix::fs::PermissionsExt;
use std::path::PathBuf;
use std::process::{Command, Stdio};

pub struct Cli {
    pub version: String,
    pub force_sudo: bool,
    pub pkg_dir: PathBuf,
}

impl Cli {
    pub fn execute(&self, args: Vec<String>) {
        if matches!(args.first(), Some(arg) if arg == "version") {
            println!("{}", self.version);
            return;
        }

        if self.force_sudo && env::var_os("SUDO_COMMAND").is_none() {
            logger::fatalf(format_args!("dyn must be run as a super user.\n"));
        }

        let mut act = match args.get(0) {
            Some(act) if is_action(act) => act.clone(),
            _ => logger::fatalf(format_args!(
                "valid actions: update|install|remove|fetch.\n"
            )),
        };

        let mut pkg_arg_index = 1;

        if act == "fetch" {
            self.fetch();
            if let Some(next) = args.get(pkg_arg_index) {
                if !matches_action_after_fetch(next) {
                    return;
                }
                act = next.clone();
                pkg_arg_index += 1;
            } else {
                return;
            }
        }

        let pkg = match args.get(pkg_arg_index) {
            Some(pkg) => pkg.clone(),
            None if act == "update" => "dyn".to_string(),
            None => logger::fatalf(format_args!(
                "please specify the package you wish to {}.\n",
                act
            )),
        };

        self.execute_package(&pkg, &act);
        logger::dynf(format_args!(
            "if you wish to contribute, make sure to check out {}\n",
            hyperlink("our github page", "https://github.com/restartfu/dyn")
        ));
        logger::dynf(format_args!("done {} package {}.\n", verb(&act), pkg));
    }

    fn execute_package(&self, pkg: &str, act: &str) {
        let target_pkg_path = self.pkg_dir.join(pkg);
        if !target_pkg_path.exists() {
            logger::fatalf(format_args!(
                "no package found with the name {}, maybe run 'dyn fetch', and try again?\n",
                pkg
            ));
        }

        let script_path = target_pkg_path.join("DYNPKG");
        if !script_path.exists() {
            logger::fatalf(format_args!(
                "no DYNPKG file found for package {}, maybe run 'dyn fetch', and try again?\n",
                pkg
            ));
        }

        let script_buf = fs::read_to_string(&script_path).unwrap_or_else(|_| {
            logger::fatalf(format_args!(
                "could not read DYNPKG file for package {}",
                pkg
            ));
        });

        let script = format!(
            "{script}\n{act}\nif [ -n \"$maintainers\" ]; then\n\tcredits=$(echo \"special thanks to ( $maintainers ) for maintaining this package\")\n\techo $credits;\n fi\n",
            script = script_buf,
            act = act
        );

        let tmp_script_path = env::temp_dir().join("dyn-pkg").join(pkg).join("script.sh");
        if let Some(tmp_dir) = tmp_script_path.parent() {
            let _ = fs::remove_dir_all(tmp_dir);
            fs::create_dir_all(tmp_dir).unwrap_or_else(|err| {
                logger::fatalf(format_args!(
                    "could not create temporary directory {}: {}\n",
                    tmp_dir.display(),
                    err
                ));
            });
        }

        fs::write(&tmp_script_path, script).unwrap_or_else(|err| {
            logger::fatalf(format_args!(
                "could not write temporary package script: {}\n",
                err
            ));
        });
        let _ = fs::set_permissions(&tmp_script_path, fs::Permissions::from_mode(0o755));

        logger::dynf(format_args!("{} package {}.\n", verb(act), pkg));

        let mut child = Command::new("sh")
            .arg(&tmp_script_path)
            .stdout(Stdio::piped())
            .stderr(Stdio::inherit())
            .spawn()
            .unwrap_or_else(|err| {
                logger::fatalf(format_args!(
                    "failed to start package script {}: {}\n",
                    tmp_script_path.display(),
                    err
                ));
            });

        if let Some(stdout) = child.stdout.take() {
            let reader = BufReader::new(stdout);
            for line in reader.lines() {
                match line {
                    Ok(line) => logger::infof(format_args!("{}\n", line)),
                    Err(err) => {
                        logger::errorf(format_args!("error reading package output: {}\n", err))
                    }
                }
            }
        }

        let status = child.wait().unwrap_or_else(|err| {
            logger::fatalf(format_args!("failed to run package script: {}\n", err));
        });
        if !status.success() {
            logger::errorf(format_args!(
                "package script exited with status {}\n",
                status
            ));
        }

        if let Some(tmp_dir) = tmp_script_path.parent() {
            let _ = fs::remove_dir_all(tmp_dir);
        }
    }

    fn fetch(&self) {
        logger::dynf(format_args!("fetching dyn-pkg git repository.\n"));

        let _ = fs::remove_dir_all(&self.pkg_dir);
        let status = Command::new("git")
            .args([
                "clone",
                "--depth",
                "1",
                "https://github.com/RestartFU/dyn-pkg",
            ])
            .arg(&self.pkg_dir)
            .stdout(Stdio::inherit())
            .stderr(Stdio::inherit())
            .status()
            .unwrap_or_else(|err| {
                logger::fatalf(format_args!("failed to launch git: {}\n", err));
            });

        if !status.success() {
            logger::fatalf(format_args!(
                "error fetching dyn package repository (exit status {}).\n",
                status
            ));
        }
    }
}

fn is_action(act: &str) -> bool {
    matches!(act, "update" | "install" | "remove" | "fetch")
}

fn matches_action_after_fetch(act: &str) -> bool {
    matches!(act, "update" | "install" | "remove")
}

fn verb(s: &str) -> &'static str {
    match s {
        "install" => "installing",
        "update" => "updating",
        "remove" => "removing",
        "fetch" => "fetching",
        _ => "working on",
    }
}

fn hyperlink(text: &str, url: &str) -> String {
    format!("\x1b]8;;{url}\x1b\\{}\x1b]8;;\x1b\\", text.yellow())
}
