<template>
	<div
		class="glass-card"
		:class="{ hoverable }"
	>
		<slot />
	</div>
</template>

<script setup lang="ts">
defineProps<{
	hoverable?: boolean;
}>();
</script>

<style scoped>
.glass-card {
	position: relative;

	/* Glass background */
	background: rgba(255, 255, 255, 0.15);
	backdrop-filter: blur(14px);
	-webkit-backdrop-filter: blur(14px);

	/* Border */
	border: 1px solid rgba(255, 255, 255, 0.25);

	/* Shadow */
	box-shadow:
		0 8px 32px rgba(0, 0, 0, 0.25),
		inset 0 1px 0 rgba(255, 255, 255, 0.3);

	/* Shape */
	border-radius: 16px;
	padding: 3rem 3.5rem;

	/* Smooth feel */
	transition:
		transform 0.25s ease,
		box-shadow 0.25s ease;
}

/* Optional hover effect */
.glass-card.hoverable:hover {
	transform: translateY(-4px);
	box-shadow:
		0 16px 40px rgba(0, 0, 0, 0.35),
		inset 0 1px 0 rgba(255, 255, 255, 0.35);
}
</style>
