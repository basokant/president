<script lang="ts">
	import { Switch } from './ui/switch';

	export let name: string;
	export let description: string;
	export let active = false;
</script>

<div class="space-y-3 py-2">
	<div class="flex items-center gap-3">
		<Switch bind:checked={active} />
		<slot />
		<h4 class="text-lg font-medium">{name}</h4>
	</div>
	<p class="font-light text-muted-foreground">
		{description}
	</p>
</div>
