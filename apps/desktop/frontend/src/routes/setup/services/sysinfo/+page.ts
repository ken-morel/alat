import { SettingsGetSysInfo } from "$lib/wails/wailsjs/go/app/App";

export const load = async () => {
  const settings = await SettingsGetSysInfo();
  return {
    settings,
  };
};
