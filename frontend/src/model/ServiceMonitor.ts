export interface ServiceMonitor {
	lastCheck: string
	lastCheckDate: Date
	httpChecks?: HTTPChecks[]
	tcpChecks?: TCPChecks[]
	dnsChecks?: DNSChecks[]
}

export interface HTTPChecks {
	displayName: string
	status: boolean
}

export interface TCPChecks {
	displayName: string
	status: boolean
}

export interface DNSChecks {
	displayName: string
	status: boolean
}