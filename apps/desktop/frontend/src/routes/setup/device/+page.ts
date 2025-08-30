import {
  GetAlatColors,
  SettingsGetDeviceColor,
  SettingsGetDeviceName,
} from "$lib/wails/wailsjs/go/app/App";
import type { PageLoad } from "./$types";

export const load: PageLoad = async (): Promise<{
  deviceName: string;
  deviceColor: string;
  alatColors: { name: string; hex: string }[];
}> => {
  const name = await SettingsGetDeviceName();
  const deviceColor = await SettingsGetDeviceColor();
  const colors = await GetAlatColors();
  return {
    deviceName: name,
    deviceColor,
    alatColors: colors.map((color) => ({
      name: color.Name,
      hex: color.Hex,
    })),
  };
};
