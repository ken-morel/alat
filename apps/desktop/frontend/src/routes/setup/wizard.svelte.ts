// This file acts as a simple, reactive store for the setup wizard's state.
// By using `$state`, any component that imports and uses these variables
// will automatically update when their values change.

export const wizardState = {
  nextUrl: $state<string | null>(null),
  prevUrl: $state<string | null>(null),
};
