import type { Actions } from './$types';
import { fail } from '@sveltejs/kit';
import { message, superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { getPlayers } from '$lib/queries/players';
import { createPlayerSchema } from '$lib/types/player';
import { getAvailableColours } from '$lib/utils/colour';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	const players = await getPlayers(params.gameId);
	const playerColour = players.filter((p) => p.isUser)[0]?.colour;
	const availableColours = getAvailableColours(players, playerColour);
	const playerSchema = createPlayerSchema(availableColours);

	const playerForm = await superValidate(zod(playerSchema));

	return {
		playerForm
	};
};

export const actions: Actions = {
	player: async ({ params, request }) => {
		const players = await getPlayers(params.gameId);
		const playerColour = players.filter((p) => p.isUser)[0]?.colour;
		const availableColours = getAvailableColours(players, playerColour);
		const playerSchema = createPlayerSchema(availableColours);

		const playerForm = await superValidate(request, zod(playerSchema));

		if (!playerForm.valid) return fail(400, { form: playerForm });

		return message(playerForm, 'player form submitted');
	}
};
