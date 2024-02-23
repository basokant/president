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

	$: user = players[0];
	$: otherPlayers = players.filter((player) => player !== user);
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
	<Button variant="filled">
		<PlusCircle />
		<span>Add Computer</span>
	</Button>
</div>
