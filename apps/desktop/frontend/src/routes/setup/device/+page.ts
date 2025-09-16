import {
  GetAlatColors,
  SettingsGetDeviceColorName,
  SettingsGetDeviceName,
} from "$lib/wails/wailsjs/go/app/App";
import type { PageLoad } from "./$types";

export const load: PageLoad = async () => {
  const name = await SettingsGetDeviceName();
  const deviceColorName = await SettingsGetDeviceColorName();
  const colors = await GetAlatColors();
  return {
    deviceName: name,
    deviceColorName,
    alatColors: colors.map((color) => ({
      name: color.name,
      hex: color.hex,
    })),
  };
};
