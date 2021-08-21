import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AppService {

  constructor(private httpClient:HttpClient) { }

  public CalculateExpression(expression:string):Observable<string> {
    return this.httpClient.post<{result:string}>("/expressions", {
      expression:expression
    }).pipe(
      map(x=>x.result)
    )
  }
}
