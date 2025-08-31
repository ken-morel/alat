import { SettingsGetFolderSyncSettings } from "$lib/wails/wailsjs/go/app/App";
import type { config } from "$lib/wails/wailsjs/go/models";

export const load = async (): Promise<{
  settings: config.FolderSyncSettings;
}> => {
  const settings = await SettingsGetFolderSyncSettings();
  return {
    settings,
  };
};
