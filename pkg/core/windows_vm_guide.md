# Minimalist Windows VM Setup for Wails Testing

This guide outlines the steps to create an extremely lightweight, disposable Windows VM for application testing, specifically targeting a minimal disk footprint.

**Goal:** A Windows VM using **< 10GB** of host disk space.
**Use Case:** Sandboxed testing for Wails applications where security and features are not a priority.
**Primary Tool:** Tiny11 (an unofficial, stripped-down Windows).

---

### 1. Get the Operating System

The best choice for this use case is **Tiny11**.

- **What it is:** An unofficial, community project by NTDEV that removes most non-essential components from Windows 11.
- **How to get it:** Search online for **"NTDEV Tiny11 Internet Archive"** to find the official ISO download page.

---

### 2. Create the VM with KVM/QEMU (virt-manager)

When setting up the virtual machine in `virt-manager`:

- **Disk Format:** Ensure you use `qcow2` (this is the default). This format supports thin provisioning, meaning the virtual disk file only uses as much space on your host as the data inside it.
- **Maximum Disk Size:** Set this to **20 GB**. The actual file will start much smaller and will likely never reach this size with the steps below.
- **RAM:** **2-4 GB** is sufficient for most testing.
- **CPU Cores:** **2 cores** is a good starting point.

---

### 3. Aggressive Post-Installation Shrinking

After you have installed Tiny11 in the VM, perform these steps inside the Windows environment to minimize its footprint.

#### a. Completely Disable Windows Update

This prevents the OS from downloading updates and growing in size.

1.  Press `Win + R`, type `services.msc`, and press Enter.
2.  Find the **Windows Update** service in the list.
3.  Double-click it.
4.  Change the "Startup type" to **Disabled**.
5.  Click the **Stop** button to halt the service if it's running.
6.  Navigate to the **Recovery** tab.
7.  Set all three failure responses ("First failure", "Second failure", "Subsequent failures") to **"Take No Action"**.
8.  Click **Apply** and **OK**.

#### b. Disable the Page File

The page file (`pagefile.sys`) is used as virtual RAM. Disabling it saves disk space equal to its size.

1.  Open the Start Menu and search for "View advanced system settings".
2.  On the **Advanced** tab, click the **Settings...** button under the "Performance" section.
3.  In the new window, go to the **Advanced** tab.
4.  Under the "Virtual memory" section, click **Change...**.
5.  Uncheck the box for **"Automatically manage paging file size for all drives"**.
6.  Select the **"No paging file"** radio button.
7.  Click **Set**, then **OK**.
8.  You will be prompted to restart the VM for the changes to take effect.

#### c. Disable Hibernation

The hibernation file (`hiberfil.sys`) reserves disk space equal to the VM's RAM.

1.  Open **Command Prompt** or **PowerShell** as an Administrator.
2.  Execute the following command:
    ```cmd
    powercfg /h off
    ```

#### d. Run CompactOS

This built-in utility compresses the core operating system files, saving significant space.

1.  Open **Command Prompt** or **PowerShell** as an Administrator.
2.  Execute the following command:
    ```cmd
    compact.exe /CompactOS:always
    ```

---

### Expected Outcome

After following these steps, your Tiny11 VM's actual file size (`.qcow2`) on your Ubuntu host should be **well under 10 GB**, providing a fast, lean environment for your Wails testing.
