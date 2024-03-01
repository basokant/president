import type { Player } from '$lib/types/player';

export async function getPlayers(gameId: string): Promise<Player[]> {
	// TODO: implement fetch for getting players from backend for specific game
	return [
		{
			id: 0,
			name: 'Kevin',
			colour: 'indigo',
			isHost: true
		},
		{
			id: 1,
			name: 'Yamaan',
			colour: 'red'
		},
		{
			id: 2,
			name: 'Asith',
			colour: 'blue'
		}
	];
}
