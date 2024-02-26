<script lang="ts">
	import { cn } from '$lib/utils';
	import { Bot, Crown } from 'lucide-svelte';
	import type { Player } from '.';
	import { fly } from 'svelte/transition';
	import AddEditPlayer from './add-edit-player.svelte';

	export let player: Player | undefined;
	export let isUser = false;
</script>

<div
	in:fly
	class={cn(
		'flex h-10 items-center py-1 pl-3',
		isUser && player?.colour && `border-2 border-${player?.colour}-500`
	)}
>
	<div class="flex flex-1 items-center gap-3">
		<div
			class={cn('size-3 rounded-full bg-background', player?.colour && `bg-${player.colour}-500`)}
		/>
		{#if player?.isHost}
			<Crown class="size-4" />
		{/if}
		{#if player?.isComputer}
			<Bot class="size-4" />
		{/if}
		<span class="font-semibold">{player?.name}</span>
	</div>
	{#if isUser}
		<AddEditPlayer {player} />
	{/if}
</div>
