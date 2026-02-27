export interface IPayment {
	id: number;
	merchantName: string;
	amount: number;
	date: string;
	status: "SUCCESS" | "PENDING" | "FAILED";
}
