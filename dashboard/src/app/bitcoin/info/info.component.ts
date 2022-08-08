import {Component, OnDestroy, OnInit} from '@angular/core';
import {Bitcoin} from "../bitcoin";
import {BitcoinService} from "../bitcoin.service";
import {Util} from "../../util";

@Component({
  selector: 'app-info',
  templateUrl: './info.component.html',
  styleUrls: ['./info.component.css']
})
export class InfoComponent implements OnInit, OnDestroy {
  model?: Bitcoin
  interval?: number

  transformBytes = Util.transformBytes

  constructor(private btcService: BitcoinService) {
    this.interval = setInterval(() => {
      this.fetchBitcoinInfo()
    }, 2000)
  }

  ngOnInit(): void {
    this.fetchBitcoinInfo()
  }

  private fetchBitcoinInfo(): void {
    this.btcService.fetchInfo().subscribe(data => {
      this.updateBitcoinInfo(data)
    }, error => {
      console.error(error)
    })
  }

  private updateBitcoinInfo(info: Bitcoin): void {
    this.model = info
  }

  ngOnDestroy() {
    if (this.interval === undefined)
      return
    clearInterval(this.interval)
  }
}
