export interface IPayment {
	id: number;
	merchant_name: string;
	amount: number;
	date: string;
	status: "SUCCESS" | "PENDING" | "FAILED";
}

export interface ISummary {
	failed: number;
	success: number;
	total: number;
}

export interface IResponsePayments {
	summary: ISummary;
	payments: IPayment[];
}
