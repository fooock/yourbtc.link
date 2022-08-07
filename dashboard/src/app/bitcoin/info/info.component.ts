import {Component, OnInit} from '@angular/core';
import {Bitcoin} from "../bitcoin";
import {BitcoinService} from "../bitcoin.service";

@Component({
  selector: 'app-info',
  templateUrl: './info.component.html',
  styleUrls: ['./info.component.css']
})
export class InfoComponent implements OnInit {
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
