import { getPlayers } from '$lib/queries/players';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async ({ params }) => {
	const players = await getPlayers(params.gameId);
	return {
		players
	};
};
