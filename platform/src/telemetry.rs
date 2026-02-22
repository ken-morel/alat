use battery::Manager;
use nlem::service::telemetry;
use sysinfo::{Disks, System};

pub async fn collect_info(info: &mut telemetry::TelemetryInfo) -> Result<(), String> {
    let mut sys = System::new_all();
    sys.refresh_all();
    tokio::time::sleep(std::time::Duration::from_millis(500)).await;

    for cpu in sys.cpus() {
        info.cpu_usages.push(cpu.cpu_usage());
    }
    info.memory_used_mb = (sys.used_memory() / (1 << 20))
        .try_into()
        .unwrap_or_default();
    info.memory_total_mb = (sys.total_memory() / (1 << 20))
        .try_into()
        .unwrap_or_default();
    info.swap_used_mb = (sys.used_swap() / (1 << 20)).try_into().unwrap_or_default();
    info.swap_total_mb = (sys.total_swap() / (1 << 20))
        .try_into()
        .unwrap_or_default();

    for disk in Disks::new_with_refreshed_list().list() {
        info.disk_names
            .push(String::from(disk.name().to_str().unwrap_or("<None>")));
        // two step division, to prevent overflows
        let total = ((disk.total_space() / 8 << 20) as f32 / (1 << 10) as f32)
            .try_into()
            .unwrap_or_default();
        info.disk_used_spaces_gb
            .push(total - ((disk.available_space() / 8 << 20) as f32 / (1 << 10) as f32));
        info.disk_total_spaces_gb.push(total);
    }

    if let Ok(batman) = Manager::new() {
        for bat in batman
            .batteries()
            .map_or_else(|_| vec![], |b| b.into_iter().flatten().collect::<Vec<_>>())
        {
            info.batteries_charging
                .push(bat.state() == battery::State::Charging);
            info.batteries_percent
                .push((bat.energy().value / bat.energy_full().value * 100.0) as u32);
        }
    }

    info.hostname = hostname::get()
        .unwrap_or("<unknown>".into())
        .into_string()
        .unwrap_or_else(|_| String::from("<None>"));

    info.uptime_secs = 10;
    info.uptime_secs = 10;

    info.kernel_version = String::from("0.1.0");

    info.os_version = String::from("Alat fake os");
    info.os_name = String::from("Alat os");

    Ok(())
}
