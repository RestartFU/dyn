use owo_colors::OwoColorize;
use std::fmt;
use std::io::{self, Write};
use std::process;

#[derive(Clone, Copy)]
pub struct LogWriter {
    level: Level,
}

impl LogWriter {
    pub fn info() -> Self {
        Self { level: Level::Info }
    }
}

impl Write for LogWriter {
    fn write(&mut self, buf: &[u8]) -> io::Result<usize> {
        let text = String::from_utf8_lossy(buf);
        log(self.level, text.as_ref());
        Ok(buf.len())
    }

    fn flush(&mut self) -> io::Result<()> {
        io::stdout().flush()
    }
}

pub fn debugf(args: fmt::Arguments<'_>) {
    log(Level::Debug, args);
}

pub fn fatalf(args: fmt::Arguments<'_>) -> ! {
    log(Level::Fatal, args);
    process::exit(0);
}

pub fn fatal(msg: &str) -> ! {
    log(Level::Fatal, msg);
    process::exit(0);
}

pub fn errorf(args: fmt::Arguments<'_>) {
    log(Level::Error, args);
}

pub fn error(msg: &str) {
    log(Level::Error, msg);
}

pub fn infof(args: fmt::Arguments<'_>) {
    log(Level::Info, args);
}

pub fn info(msg: &str) {
    log(Level::Info, msg);
}

pub fn dynf(args: fmt::Arguments<'_>) {
    log(Level::Dyn, args);
}

pub fn aqua(text: &str) -> String {
    text.cyan().to_string()
}

fn log(level: Level, msg: impl fmt::Display) {
    let (label, color) = match level {
        Level::Debug => ("DEBU", LabelColor::Blue),
        Level::Fatal => ("FATA", LabelColor::Red),
        Level::Error => ("ERRO", LabelColor::BrightRed),
        Level::Info => ("INFO", LabelColor::Yellow),
        Level::Dyn => ("DYN ", LabelColor::Cyan),
    };

    let prefix = color.paint(label);
    let divider = "|".truecolor(128, 128, 128);

    print!("{prefix}{divider} {msg}");
    io::stdout().flush().ok();
}

#[derive(Clone, Copy)]
enum Level {
    Debug,
    Fatal,
    Error,
    Info,
    Dyn,
}

#[derive(Clone, Copy)]
enum LabelColor {
    Blue,
    Red,
    BrightRed,
    Yellow,
    Cyan,
}

impl LabelColor {
    fn paint(self, text: &str) -> String {
        match self {
            LabelColor::Blue => text.blue().to_string(),
            LabelColor::Red => text.red().to_string(),
            LabelColor::BrightRed => text.bright_red().to_string(),
            LabelColor::Yellow => text.yellow().to_string(),
            LabelColor::Cyan => text.cyan().to_string(),
        }
    }
}
