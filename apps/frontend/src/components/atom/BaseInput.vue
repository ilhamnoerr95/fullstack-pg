<script setup lang="ts">
defineProps<{
	modelValue: string;
	label?: string;
	type?: string;
	placeholder?: string;
	error?: string;
}>();

defineEmits(["update:modelValue"]);
</script>

<template>
	<div class="field">
		<label v-if="label">{{ label }}</label>

		<input
			:type="type || 'text'"
			:placeholder="placeholder"
			:value="modelValue"
			@input="
				$emit('update:modelValue', ($event.target as HTMLInputElement).value)
			"
			:class="{ error: error }"
		/>

		<small
			v-if="error"
			class="error-text"
			>{{ error }}</small
		>
	</div>
</template>

<style scoped>
.field {
	display: flex;
	flex-direction: column;
	gap: 10px;
	width: 23rem;

	margin-bottom: 2rem;
}

label {
	font-size: 12px;
	color: black;
	text-align: left;
}

input {
	/* width: 100%; */
	padding: 10px 12px;
	border-radius: 6px;
	border: 1px solid #ccc;
	background: #fff;
	color: black;
}

input.error {
	border-color: #e74c3c;
}

.error-text {
	font-size: 12px;
	color: #e74c3c;
}
</style>
