import type { PageLoad } from "./$types";
import {
  GetFoundDevices,
  GetConnectedDevices,
} from "$lib/wails/wailsjs/go/app/App";

export const load: PageLoad = async () => {
  const foundDevices = (await GetFoundDevices()) || [];
  const connectedDevices = await GetConnectedDevices();
  const devices = [];
  for (const foundDev of foundDevices)
    for (const connDev of connectedDevices)
      if (connDev.Info.ID !== foundDev.Info.ID) devices.push(foundDev);

  return {
    found: devices,
  };
};
