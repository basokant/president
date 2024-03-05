import { z } from 'zod';

export type Player = {
	id: number;
	name: string;
	isUser?: boolean;
	isHost?: boolean;
	isComputer?: boolean;
	colour: Colour;
};

export const colours = [
	'red',
	'orange',
	'yellow',
	'lime',
	'emerald',
	'cyan',
	'blue',
	'indigo',
	'purple',
	'pink'
] as const;

export type Colour = (typeof colours)[number];

export function createPlayerSchema(
	availableColours: Colour[],
	initialName?: string,
	initialColour?: Colour
) {
	let nameSchema = z.string().min(3).max(15);
	let colourSchema = z.enum(colours).refine((c) => availableColours.includes(c));

	// if (!!initialName && !!initialColour) {
	// 	const playerSchema = z.object({
	// 		name: nameSchema.default(initialName),
	// 		colour: colourSchema.default(initialColour)
	// 	});
	// 	return playerSchema;
	// }
	//
	// if (availableColours.length >= 0 && !!availableColours[0]) {
	// 	const playerSchema = z.object({
	// 		name: nameSchema,
	// 		colour: colourSchema.default(availableColours[0])
	// 	});
	// 	return playerSchema;
	// }

	const playerSchema = z.object({
		name: nameSchema,
		colour: colourSchema
	});

	return playerSchema;
}

export const gameSchema = z.object({
	name: z.string().min(3).max(15),
	colour: z.enum(colours)
});
