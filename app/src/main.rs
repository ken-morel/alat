
mod ui;

use slint::ComponentHandle;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let window = ui::MainWindow::new()?;

    window.run()?;

    Ok(())
}
