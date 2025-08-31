import { ConfigReady } from "$lib/wails/wailsjs/go/app/App";
import { redirect } from "@sveltejs/kit";

export const load = async () => {
  const ready = await ConfigReady();
  throw redirect(300, ready ? "/dashboard" : "/setup");
};
