import { writable } from 'svelte/store';

export type BackendConfig = {
	folderPath: string;
	logLevel?: string;
	// add other config fields as needed
};

export const backendConfig = writable<BackendConfig | null>(null);
