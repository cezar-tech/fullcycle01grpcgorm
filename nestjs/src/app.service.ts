import { Metadata } from '@grpc/grpc-js';
import { Injectable } from '@nestjs/common';
import { Client, ClientGrpc, Transport } from '@nestjs/microservices';
import { Observable } from 'rxjs';
import { ProductDto } from './dto/product.dto';

interface ProductGrpcClient {
  CreateProduct(
    data: {
      name: string;
      description: string;
      price: number;
    },
    metadata?: Metadata,
  ): Observable<object>;
  FindProducts(
    {},
    metadata?: Metadata,
  ): Observable<
    [
      {
        name: string;
        description: string;
        price: number;
        id: number;
      },
    ]
  >;
}

@Injectable()
export class AppService {
  @Client({
    transport: Transport.GRPC,
    options: {
      url: 'go:50051',
      package: 'github.com.codeedu.codepix',
      protoPath: '../proto/product.proto',
      loader: { keepCase: true },
    },
  })
  clientGrpc: ClientGrpc;

  getProducts(): Observable<
    [{ name: string; description: string; price: number; id: number }]
  > {
    const svc = this.clientGrpc.getService<ProductGrpcClient>('ProductService');
    return svc.FindProducts({});
  }

  createProduct(createDto: ProductDto): Observable<object> {
    const svc = this.clientGrpc.getService<ProductGrpcClient>('ProductService');
    console.log(createDto);
    return svc.CreateProduct({
      description: createDto.description,
      price: createDto.price,
      name: createDto.name,
    });
  }
}
