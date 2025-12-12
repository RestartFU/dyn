mod cli;
mod logger;

use clap::Parser;
use std::path::PathBuf;

const VERSION: &str = "v0.1.4";

#[derive(Parser, Debug)]
#[command(version = None, about = "Dyn: The Dynamic Linux Package Manager", trailing_var_arg = true)]
struct Args {
    /// Allow running without sudo privileges.
    #[arg(long, default_value_t = false)]
    nosudo: bool,
    /// Path where dyn packages are stored.
    #[arg(long, default_value = "/usr/local/dyn-pkg")]
    pkgdir: String,
    /// Action and package arguments.
    #[arg()]
    rest: Vec<String>,
}

fn main() {
    let args = Args::parse();

    let cli = cli::Cli {
        version: logger::aqua(VERSION),
        force_sudo: !args.nosudo,
        pkg_dir: PathBuf::from(args.pkgdir),
    };

    cli.execute(args.rest);
}
