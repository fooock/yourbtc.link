export class Bitcoin {
  constructor(
    public BlockchainInfo: BlockchainInfo,
    public NetworkInfo: NetworkInfo,
    public NetTotals: NetTotals
  ) {
  }
}

export class BlockchainInfo {
  constructor(
    public verificationprogress: number,
    public blocks: number,
    public headers: number,
    public bestblockhash: string,
    public size_on_disk: number
  ) {
  }
}

export class NetworkInfo {
  constructor(
    public subversion: string,
    public connections: number,
    public connections_in: number,
    public connections_out: number
  ) {
  }
}

export class NetTotals {
  constructor(
    public totalbytesrecv: number,
    public totalbytessent: number
  ) {
  }
}
