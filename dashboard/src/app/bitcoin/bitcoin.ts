export class Bitcoin {
  constructor(
    public BlockchainInfo: BlockchainInfo,
    public NetworkInfo: NetworkInfo,
    public NetTotals: NetTotals,
    public ChainTxStats: ChainTxStats
  ) {
  }
}

export class BlockchainInfo {
  constructor(
    public verificationprogress: number,
    public blocks: number,
    public headers: number,
    public bestblockhash: string,
    public size_on_disk: number,
    public chain: string,
    public chainwork: string,
    public difficulty: number
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

export class ChainTxStats {
  constructor(
    public txcount: number,
    public txrate: number
  ) {
  }
}
