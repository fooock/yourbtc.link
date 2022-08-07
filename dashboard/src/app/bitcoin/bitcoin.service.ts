import {Injectable} from '@angular/core';
import {environment} from "../../environments/environment";
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {Bitcoin} from "./bitcoin";

@Injectable({
  providedIn: 'root'
})
export class BitcoinService {
  private url: string = `${environment.apiUrl}/api/v1`

  constructor(private client: HttpClient) {
  }

  fetchInfo(): Observable<Bitcoin> {
    const path = `${this.url}/info`
    return this.client.get<Bitcoin>(path)
  }
}

