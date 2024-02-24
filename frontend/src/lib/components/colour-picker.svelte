<script lang="ts">
	import { cn } from '$lib/utils';
	import { colours } from './players';
	import * as ToggleGroup from './ui/toggle-group';
	import { ToggleGroup as ToggleGroupPrimitive } from 'bits-ui';

	export let selectedColour: string = colours[0];
</script>

<ToggleGroup.Root
	type="single"
	class="flex flex-wrap justify-between gap-2"
	bind:value={selectedColour}
>
	{#each colours as colour, i (i)}
		<ToggleGroupPrimitive.Item value={colour} aria-label={`Select ${colour}`} asChild let:builder>
			<div
				use:builder.action
				{...builder}
				class={cn(
					'size-8 rounded-full',
					`bg-${colour}-500`,
					selectedColour === colour &&
						` ring-2 ring-offset-4 ring-offset-background ring-${colour}-500`
				)}
			/>
		</ToggleGroupPrimitive.Item>
	{/each}
</ToggleGroup.Root>
