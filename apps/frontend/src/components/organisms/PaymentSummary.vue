<script setup lang="ts">
import { computed } from "vue";
import { IPayment } from "../../types/payment";

const props = defineProps<{
	payments: IPayment[];
}>();

const total = computed(() => props.payments.length);

const success = computed(
	() => props.payments.filter((p) => p.status === "SUCCESS").length,
);

const failed = computed(
	() => props.payments.filter((p) => p.status === "FAILED").length,
);
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
