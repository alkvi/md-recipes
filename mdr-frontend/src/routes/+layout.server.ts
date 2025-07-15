import type { LayoutServerLoad } from './$types';
import type { BackendConfig } from '$lib/stores/backendConfig';

export const load: LayoutServerLoad = async ({ fetch }) => {
	const res = await fetch('http://localhost:3000/config');
	const config: BackendConfig = await res.json();
	return { config };
};
