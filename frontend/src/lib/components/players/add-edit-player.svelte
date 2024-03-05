<script lang="ts">
	import { LogIn, Pencil, Save } from 'lucide-svelte';
	import { Button } from '../ui/button';
	import * as Dialog from '../ui/dialog';
	import { Input } from '../ui/input';
	import ColourPicker from '../colour-picker.svelte';
	import { gameSchema } from '$lib/types/player';
	import { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import * as Form from '$lib/components/ui/form';
	import { page } from '$app/stores';
	import Label from '../ui/label/label.svelte';

	const data: SuperValidated<Infer<typeof gameSchema>> = $page.data.playerForm;

	const form = superForm(data);

	const { form: formData, enhance } = form;

	export let open = false;
	export let action = '?/player';
	export let isEdit = false;
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
				<Form.Control let:attrs>
					<Form.Label>What's your Nickname?</Form.Label>
					<Input {...attrs} required autofocus type="text" bind:value={$formData.name} />
				</Form.Control>
			</Form.Field>
			<Form.Fieldset {form} name="colour" class="space-y-2 py-1">
				<Form.Legend>Choose a Colour!</Form.Legend>
				<ColourPicker bind:value={$formData.colour} />
			</Form.Fieldset>
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
