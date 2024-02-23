<script lang="ts">
	import { cn } from '$lib/utils';
	import { Bot, Crown, Pencil } from 'lucide-svelte';
	import type { Player } from '.';
	import { Button } from '../ui/button';

	export let player: Player;
	export let isUser = false;
</script>

<div
	class={cn(
		'flex h-10 items-center pl-3',
		isUser && player.colour && `border-2 border-${player.colour}-500`
	)}
>
	<div class="flex flex-1 items-center gap-3">
		<div
			class={cn('size-3 rounded-full bg-background', player.colour && `bg-${player.colour}-500`)}
		/>
		{#if player.isHost}
			<Crown class="size-4" />
		{/if}
		{#if player.isComputer}
			<Bot class="size-4" />
		{/if}
		<span class="font-semibold">{player.name}</span>
	</div>
	{#if isUser}
		<Button class="size-9 p-1" variant="ghost">
			<Pencil class="size-4" />
		</Button>
	{/if}
</div>
