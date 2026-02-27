<script setup lang="ts">
import { computed } from "vue";
import { IPayment } from "../../types/payment";

const props = defineProps<{
	summary: {
		failed: number;
		success: number;
		total: number;
	};
}>();

const total = computed(() => props.summary.total);

const success = computed(() => props.summary.success);
const failed = computed(() => props.summary.failed);
</script>

<template>
	<a-card style="margin-bottom: 20px">
		<a-space size="large">
			<span><strong>Total:</strong> {{ total }}</span>
			<span style="color: green">
				<strong>Success:</strong> {{ success }}
			</span>
			<span style="color: red"> <strong>Failed:</strong> {{ failed }} </span>
		</a-space>
	</a-card>
</template>
