<script lang="ts">
	import { cn } from '$lib/utils';
	import { Bot, Crown, LogIn, Pencil } from 'lucide-svelte';
	import type { Player } from '.';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import ColourPicker from '$lib/components/colour-picker.svelte';
	import { fly } from 'svelte/transition';
	import * as Dialog from '$lib/components/ui/dialog';

	export let player: Player;
	export let isUser = false;
	export let open = true;

	function addEditPlayer() {
		open = false;
		if (player) {
			// TODO: finish edit player logic
			return;
		}
		// TODO: finish add player logic
	}
</script>

<div
	in:fly
	class={cn(
		'flex h-10 items-center py-1 pl-3',
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
		<Dialog.Root bind:open closeOnOutsideClick={false}>
			<Dialog.Trigger>
				<Button class="size-9 p-1" variant="ghost">
					<Pencil class="size-4" />
				</Button>
			</Dialog.Trigger>
			<Dialog.Content class="flex flex-col gap-6">
				<Dialog.Header>
					<Dialog.Title>Create your Player</Dialog.Title>
				</Dialog.Header>
				<div class="space-y-2 py-1">
					<h3>What's your Nickname?</h3>
					<Input type="text" />
				</div>
				<div class="space-y-2 py-1">
					<h3>Choose a Colour!</h3>
					<ColourPicker />
				</div>
				<Button class="w-44 self-end" on:click={addEditPlayer}>
					<LogIn />
					<span>Join Game</span>
				</Button>
			</Dialog.Content>
		</Dialog.Root>
	{/if}
</div>
