import type { Player } from '$lib/types/player';
import { getContext, setContext } from 'svelte';
import { derived, writable, type Readable, type Writable } from 'svelte/store';
type PlayersStore = Writable<Player[]>;

const PLAYERS_CTX = 'PLAYERS_CTX';

export function setPlayersStore(initialPlayers: Player[]): PlayersStore {
	console.log('setting PlayersStore');
	let playersStore = getPlayersStore();

	if (playersStore === undefined) {
		playersStore = writable(initialPlayers);
		setContext(PLAYERS_CTX, playersStore);
	} else {
		playersStore.set(initialPlayers);
	}

	setUserStore(playersStore);

	return playersStore;
}

export function getPlayersStore(): PlayersStore | undefined {
	console.log('getting PlayersStore');
	return getContext<PlayersStore>(PLAYERS_CTX);
}

type UserStore = Readable<Player | undefined>;

const USER_CTX = 'USER_CTX';

export function setUserStore(playersStore: PlayersStore): UserStore {
	let userStore = getUserStore();
	if (userStore === undefined) {
		userStore = derived(playersStore, ($playersStore) => $playersStore.filter((p) => p.isUser)[0]);
		setContext(USER_CTX, userStore);
	}

	return userStore;
}

export function getUserStore(): UserStore {
	console.log('getting UserStore');
	return getContext<Readable<Player | undefined>>(USER_CTX);
}
