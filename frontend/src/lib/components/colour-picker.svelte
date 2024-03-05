<script lang="ts">
	import { getPlayersStore } from '$lib/stores/players.store';
	import { cn } from '$lib/utils';
	import { type Colour } from '$lib/types/player';
	import { getAvailableColours } from '$lib/utils/colour';
	import { RadioGroup } from 'bits-ui';
	import type { ControlProps } from 'formsnap';
	import * as Form from './ui/form';
	import Label from './ui/label/label.svelte';

	export let value: Colour;
	const players = getPlayersStore();
	const availableColours = getAvailableColours($players ?? [], value);

	$: console.log(value);
</script>

<RadioGroup.Root
	orientation="horizontal"
	class="flex flex-wrap justify-between gap-4 md:gap-2"
	bind:value
>
	{#each availableColours as col, i (i)}
		<Form.Control let:attrs>
			<Form.Label hidden>Select {col}</Form.Label>
			<RadioGroup.Item value={col} asChild let:builder {...attrs}>
				<div
					use:builder.action
					{...builder}
					class={cn(
						'size-11 cursor-pointer rounded-full md:size-8',
						`bg-${col}-500`,
						value === col && `ring-2 ring-offset-4 ring-offset-background ring-${col}-500`
					)}
				/>
			</RadioGroup.Item>
		</Form.Control>
	{/each}
	<RadioGroup.Input name="colour" />
</RadioGroup.Root>
