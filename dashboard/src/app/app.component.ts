import {Component, OnInit} from '@angular/core';
import {BitcoinService} from './services/bitcoin.service'
import {Bitcoin} from "./model/bitcoin";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  model: Bitcoin | undefined

  constructor(private btcService: BitcoinService) {
  }

  ngOnInit(): void {
    this.btcService.fetchInfo()
      .subscribe(data => {
        this.model = data
      }, error => {
        console.error(error)
      })
  }
}
