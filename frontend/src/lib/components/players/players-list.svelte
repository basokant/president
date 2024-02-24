<script lang="ts">
	import { cn } from '$lib/utils';
	import { ChevronsUpDown, Construction, PlusCircle, UsersRound } from 'lucide-svelte';
	import { colours, type Player } from '.';
	import { Button } from '../ui/button';
	import PlayerItem from './player-item.svelte';
	import * as Collapsible from '$lib/components/ui/collapsible';

	let players: Player[] = [
		{
			id: 0,
			name: 'Ben',
			colour: 'emerald',
			isHost: true
		}
	];

	$: nextPlayerId = Math.max(...players.map((p) => p.id)) + 1;
	$: user = players[0];
	$: otherPlayers = players.filter((player) => player !== user);
	$: numComputers = players.filter((p) => p.isComputer).length;
	$: availableColours = colours.filter((c) => !players.some((p, i, a) => p.colour === c));

	function addComputer() {
		if (players.length == 8) {
			return;
		}

		const newComputer: Player = {
			id: nextPlayerId,
			name: `Computer ${numComputers + 1}`,
			colour: availableColours[Math.floor(Math.random() * availableColours.length)],
			isComputer: true
		};
		players = [...players, newComputer];
	}
</script>

<Collapsible.Root open class={cn('flex flex-col gap-2', $$props.class)}>
	<div class="flex flex-row items-center gap-2">
		<UsersRound strokeWidth={2} />
		<h4 class="flex-1 text-lg font-medium">
			{players.length}/8 Players
		</h4>
		<Collapsible.Trigger asChild let:builder>
			<Button builders={[builder]} variant="ghost" size="sm" class="w-9 p-0">
				<ChevronsUpDown class="h-4 w-4" />
				<span class="sr-only">Toggle</span>
			</Button>
		</Collapsible.Trigger>
	</div>
	<PlayerItem player={user} isUser />
	<Collapsible.Content class="py-0">
		{#each otherPlayers as player, i (i)}
			<PlayerItem {player} />
		{/each}
		{#if otherPlayers.length === 0}
			<div class="flex items-center gap-2 p-3">
				<Construction />
				<span>No other players</span>
			</div>
		{/if}
	</Collapsible.Content>
	{#if user.isHost}
		<Button disabled={players.length === 8} variant="filled" on:click={addComputer}>
			<PlusCircle />
			<span>Add Computer</span>
		</Button>
	{/if}
</Collapsible.Root>
