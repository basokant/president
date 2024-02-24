import Root from './button.svelte';
import { tv, type VariantProps } from 'tailwind-variants';
import type { Button as ButtonPrimitive } from 'bits-ui';

const buttonVariants = tv({
	base: 'inline-flex items-center justify-center rounded-md:w font-medium whitespace-nowrap ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50',
	variants: {
		variant: {
			default: 'bg-primary text-primary-foreground hover:bg-primary/85',
			destructive: 'bg-destructive text-destructive-foreground hover:bg-destructive/90',
			outline: 'border-2 border-input bg-background hover:bg-accent hover:text-accent-foreground',
			outlinePrimary: 'border-primary border-2 bg-background hover:bg-primary/10 text-primary',
			secondary: 'bg-secondary text-secondary-foreground hover:bg-secondary/80',
			filled: 'bg-foreground text-background hover:bg-foreground/85',
			ghost: 'hover:bg-accent hover:text-accent-foreground',
			link: 'underline-offset-4 hover:underline'
		},
		size: {
			default: 'h-10 px-3 py-2 space-x-3',
			sm: 'h-9 rounded-md px-3 space-x-3',
			lg: 'h-11 rounded-md px-8 space-x-4',
			icon: 'h-10 w-10'
		}
	},
	defaultVariants: {
		variant: 'default',
		size: 'default'
	}
});

type Variant = VariantProps<typeof buttonVariants>['variant'];
type Size = VariantProps<typeof buttonVariants>['size'];

type Props = ButtonPrimitive.Props & {
	variant?: Variant;
	size?: Size;
};

type Events = ButtonPrimitive.Events;

export {
	Root,
	type Props,
	type Events,
	//
	Root as Button,
	type Props as ButtonProps,
	type Events as ButtonEvents,
	buttonVariants
};
