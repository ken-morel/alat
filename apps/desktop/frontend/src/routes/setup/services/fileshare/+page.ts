import { SettingsGetFileSharingSettings } from "$lib/wails/wailsjs/go/app/App";
import type { config } from "$lib/wails/wailsjs/go/models";

export const load = async (): Promise<{
  settings: config.FileSharingSettings;
}> => {
  const settings = await SettingsGetFileSharingSettings();
  return {
    settings,
  };
};
