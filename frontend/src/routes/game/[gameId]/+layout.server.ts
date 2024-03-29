import { getPlayers } from '$lib/queries/players';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ params }) => {
	const players = await getPlayers(params.gameId);

	return {
		players
	};
};
