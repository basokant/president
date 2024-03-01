<script lang="ts">
	import { LogIn, Pencil, Save } from 'lucide-svelte';
	import { Button } from '../ui/button';
	import * as Dialog from '../ui/dialog';
	import { Input } from '../ui/input';
	import ColourPicker from '../colour-picker.svelte';
	import { gameSchema } from '$lib/types/player';
	import { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import * as Form from '$lib/components/ui/form';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { page } from '$app/stores';

	const data: SuperValidated<Infer<typeof gameSchema>> = $page.data.playerForm;

	const form = superForm(data, {
		validators: zodClient(gameSchema)
	});

	const { form: formData, enhance } = form;

	export let open = true;
	export let action: string = '?/player';

	const isEdit = !!$formData.name;
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
		<form class="flex flex-col gap-6" method="post" {action} use:enhance>
			<Form.Field {form} name="name" class="space-y-2 py-1">
				<h3>What's your Nickname?</h3>
				<Input required autofocus type="text" bind:value={$formData.name} />
			</Form.Field>
			<Form.Field {form} name="colour" class="space-y-2 py-1">
				<h3>Choose a Colour!</h3>
				<ColourPicker initialColour={$formData.colour} bind:colour={$formData.colour} />
			</Form.Field>
			<Form.Button class="w-44 gap-2 self-end">
				{#if isEdit}
					<Save />
					<span>Save Player</span>
				{:else}
					<LogIn />
					Join Game
				{/if}
			</Form.Button>
		</form>
	</Dialog.Content>
</Dialog.Root>
