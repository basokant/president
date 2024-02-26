<script lang="ts">
	import { LogIn, Pencil, Save } from 'lucide-svelte';
	import { Button } from '../ui/button';
	import * as Dialog from '../ui/dialog';
	import { Input } from '../ui/input';
	import ColourPicker from '../colour-picker.svelte';
	import { type Player } from '.';
	import { getPlayersStore } from '$lib/stores/players.store';
	import { toast } from 'svelte-sonner';

	let players = getPlayersStore();
	export let player: Player | undefined = undefined;
	export let open = false;

	$: isEdit = !!player;

	let name = player?.name || '';
	let colour = player?.colour;

	function addEditPlayer(): void {
		if (!name) {
			toast.error('Name is required to create a player');
			return;
		}

		if (!colour) {
			toast.error('Colour is required to create a player');
			return;
		}

		if (isEdit) {
			players.editPlayer(player?.id ?? -1, { name, colour });
			open = false;
			return;
		}

		if ($players.length === 8) {
			toast.error('No more space for players, 8 is the maximum');
		}

		let isHost = false;
		if ($players.length === 0) {
			isHost = true;
		}

		players.addPlayer({ name, colour, isUser: true, isHost });
		open = false;
	}
</script>

<Dialog.Root bind:open closeOnOutsideClick={isEdit} closeOnEscape={isEdit}>
	{#if isEdit}
		<Dialog.Trigger asChild let:builder>
			<Button builders={[builder]} {...builder} class="size-9 p-1" variant="ghost">
				<Pencil class="size-4" />
			</Button>
		</Dialog.Trigger>
	{/if}
	<Dialog.Content class="flex flex-col gap-6" closeButton={isEdit}>
		<Dialog.Header>
			<Dialog.Title>Create your Player</Dialog.Title>
		</Dialog.Header>
		<form on:submit={addEditPlayer} class="flex flex-col gap-6">
			<div class="space-y-2 py-1">
				<h3>What's your Nickname?</h3>
				<Input required autofocus type="text" bind:value={name} />
			</div>
			<div class="space-y-2 py-1">
				<h3>Choose a Colour!</h3>
				<ColourPicker bind:colour initialColour={colour} />
			</div>
			<Button class="w-44 gap-2 self-end" type="submit">
				{#if isEdit}
					<Save />
					<span>Save Player</span>
				{:else}
					<LogIn />
					Join Game
				{/if}
			</Button>
		</form>
	</Dialog.Content>
</Dialog.Root>
