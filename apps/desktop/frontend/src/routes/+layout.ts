import { ConfigReady } from "$lib/wails/wailsjs/go/app/App";
import { redirect } from "@sveltejs/kit";

export const prerender = true;
export const ssr = false;

export const load = async ({ url }: { url: URL }) => {
  const ready = await ConfigReady();
  if (!ready && !url.pathname.startsWith("/setup")) {
    throw redirect(300, "/setup");
  }
};

