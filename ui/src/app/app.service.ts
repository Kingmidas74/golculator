import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, map } from 'rxjs/operators';
import { Observable, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AppService {

  constructor(private httpClient:HttpClient) { }

  public CalculateExpression(expression:string):Observable<string> {
    return this.httpClient.post<{result:string}>("/expressions", {
      expression:expression
    }).pipe(
      map(x=>x.result),
      catchError(err=>throwError("Wrong"))
    )
  }



  public GetOperations():Observable<any[]> {
    return this.httpClient.get<any[]>("/operations")
  }



  public SaveOperation(operation:any):Observable<any> {
    return this.httpClient.post("/operations",operation)
  }
}
