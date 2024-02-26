import type { Player } from '$lib/components/players';
import { getContext, setContext } from 'svelte';
import { writable } from 'svelte/store';

function createPlayersStore(initialPlayers: Player[]) {
	const { subscribe, update } = writable<Player[]>(initialPlayers);

	function addPlayer(player: Omit<Player, 'id'>) {
		update((players) => {
			const nextPlayerId = Math.max(...players.map((p) => p.id)) + 1;
			const newPlayer = {
				id: nextPlayerId,
				...player
			};

			players.push(newPlayer);
			return players;
		});
	}

	function editPlayer(id: number, updatePlayer: Partial<Omit<Player, 'id'>>) {
		update((players) => {
			const i = players.findIndex((p) => p.id === id);
			const updatedPlayer = {
				...players[i],
				...updatePlayer
			};

			players[i] = updatedPlayer;
			return players;
		});
	}

	return {
		subscribe,
		addPlayer,
		editPlayer
	};
}

type PlayersStore = ReturnType<typeof createPlayersStore>;

const PLAYERS_CTX = 'PLAYERS_CTX';

export function setPlayersStore(initialPlayers: Player[]): PlayersStore {
	const playersStore = createPlayersStore(initialPlayers);
	setContext(PLAYERS_CTX, playersStore);
	return playersStore;
}

export function getPlayersStore(): PlayersStore {
	return getContext<PlayersStore>(PLAYERS_CTX);
}
