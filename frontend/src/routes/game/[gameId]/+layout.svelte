<script lang="ts">
	import { setPlayersStore, setUserStore } from '$lib/stores/players.store';
	import type { PageData } from './$types';

	export let data: PageData;
	$: players = setPlayersStore(data.players);
	$: user = setUserStore(players);
</script>

<slot />
