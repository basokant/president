import type { Player } from '$lib/types/player';

const players: Player[] = [
	{
		id: 0,
		name: 'Kevin',
		colour: 'indigo'
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

export async function getPlayers(gameId: string): Promise<Player[]> {
	// TODO: implement API call to Go backend for getting players from backend for specific gameId
	return players;
}

export async function upsertPlayer(
	playerData: Omit<Player, 'id'>,
	playerId?: number
): Promise<boolean> {
	// TODO: implement API call to Go backend to upsert the player.
	const isEdit = !!playerId;

	if (isEdit) {
		return await editPlayer({
			...playerData,
			id: playerId
		});
	}

	return await createPlayer(playerData);
}

async function createPlayer(newPlayerData: Omit<Player, 'id'>) {
	console.log('Creating Player', { player: newPlayerData });
	const newPlayerId = Math.max(...players.map((p) => p.id)) + 1;
	players.push({
		id: newPlayerId,
		isUser: true,
		...newPlayerData
	});

	return true;
}

async function editPlayer(updatedPlayer: Player): Promise<boolean> {
	console.log('Editing Player', { player: updatedPlayer });
	const index = players.findIndex((p) => p.id === updatedPlayer.id);
	if (index < 0) {
		return false;
	}

	players[index] = updatedPlayer;
	return true;
}
