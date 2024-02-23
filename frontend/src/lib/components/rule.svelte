<script lang="ts">
	import { Label } from './ui/label';
	import { Switch } from './ui/switch';

	export let id: string;
	export let name: string;
	export let description: string;
	export let checked = false;
</script>

<div class="space-y-3 py-2">
	<div class="flex items-center gap-3">
		<Switch {id} bind:checked />
		<slot />
		<Label for={id} class="text-lg font-medium">{name}</Label>
	</div>
	<p class="font-light text-muted-foreground">
		{description}
	</p>
</div>
