import { writable } from "svelte/store";

export const nextUrl = writable<string | null>(null);
export const prevUrl = writable<string | null>(null);
