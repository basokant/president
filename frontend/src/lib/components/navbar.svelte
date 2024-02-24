<script>
	import { Club, Diamond, Heart, Spade } from 'lucide-svelte';
	import { Button } from '$lib/components/ui/button';
	import Separator from './ui/separator/separator.svelte';
	import ModeSwitcher from '$lib/components/mode-switcher.svelte';
	import Help from '$lib/components/help.svelte';
</script>

<nav class="flex flex-row items-center justify-between py-2 pt-3">
	<Button class="px-0 text-2xl font-bold" variant="link" href="/">President</Button>
	<div class="hidden w-72 justify-between md:flex">
		<Heart strokeWidth={3} size={20} class="text-primary" />
		<Club strokeWidth={3} size={20} />
		<Diamond strokeWidth={3} size={20} class=" text-primary " />
		<Spade strokeWidth={3} size={20} />
	</div>
	<div class="flex items-center gap-2">
		<Help />
		<ModeSwitcher />
	</div>
</nav>
<Separator orientation="horizontal" />
