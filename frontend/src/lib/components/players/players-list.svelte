<script lang="ts">
	import { cn } from '$lib/utils';
	import { PlusCircle, UsersRound } from 'lucide-svelte';
	import { type Player } from '.';
	import { Button } from '../ui/button';
	import PlayerItem from './player-item.svelte';

	let players: Player[] = [
		{
			id: 0,
			name: 'Ben',
			colour: 'green-600'
		}
	];

	$: nextPlayerId = Math.max(...players.map((p) => p.id)) + 1;
	$: user = players[0];
	$: otherPlayers = players.filter((player) => player !== user);
	$: numComputers = players.filter((p) => p.isComputer).length;

	function addComputer() {
		if (players.length == 8) {
			return;
		}

		const newComputer: Player = {
			id: nextPlayerId,
			name: `Computer ${numComputers + 1}`,
			colour: 'green-600',
			isComputer: true
		};
		players = [...players, newComputer];
	}
</script>

<div class={cn('flex flex-col gap-2', $$props.class)}>
	<div class="flex flex-row items-center gap-2">
		<UsersRound strokeWidth={2} />
		<h4 class="text-lg font-medium">
			{players.length}/8 Players
		</h4>
	</div>
	<div class="py-1">
		<PlayerItem player={user} isUser />
		{#each otherPlayers as player, i (i)}
			<PlayerItem {player} />
		{/each}
	</div>
	<Button variant="filled" on:click={addComputer}>
		<PlusCircle />
		<span>Add Computer</span>
	</Button>
</div>
