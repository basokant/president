<script lang="ts">
	import { getPlayersStore } from '$lib/stores/players.store';
	import { cn } from '$lib/utils';
	import { colours, type Colour, type Player } from './players';
	import { RadioGroup } from 'bits-ui';

	export let colour: Colour = colours[0];
	export let initialColour: Colour | undefined = undefined;

	let players = getPlayersStore();

	function getAvailableColours(players: Player[]) {
		const usedColours = new Set<string>();
		players.forEach((p) => usedColours.add(p.colour));
		return colours.filter((c) => !usedColours.has(c) || c === initialColour);
	}
	$: availableColours = getAvailableColours($players);
</script>

<RadioGroup.Root
	orientation="horizontal"
	class="flex flex-wrap justify-between gap-4 md:gap-2"
	bind:value={colour}
>
	{#each availableColours as col, i (i)}
		<RadioGroup.Item value={col} aria-label={`Select ${col}`} asChild let:builder>
			<div
				use:builder.action
				{...builder}
				class={cn(
					'size-11 cursor-pointer rounded-full md:size-8',
					`bg-${col}-500`,
					colour === col && `ring-2 ring-offset-4 ring-offset-background ring-${col}-500`
				)}
			/>
		</RadioGroup.Item>
	{/each}
</RadioGroup.Root>
