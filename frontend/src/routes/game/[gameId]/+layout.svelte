<script lang="ts">
	import { setPlayersStore } from '$lib/stores/players.store';
	import type { PageData } from './$types';

	export let data: PageData;
	const players = setPlayersStore(data.players);
</script>

<slot />
