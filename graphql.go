package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var SensorHistory = graphql.NewObject(graphql.ObjectConfig{
	Name: "HistoryEvent",
	Fields: graphql.Fields{
		"time": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if h, ok := p.Source.(tempSensorHistory); ok {
					return h.Time, nil
				}
				return nil, nil
			},
		},
		"value": &graphql.Field{
			Type: graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if h, ok := p.Source.(tempSensorHistory); ok {
					return h.Value, nil
				}
				return nil, nil
			},
		},
	},
})

// Sensor is graphql type for sensor
var Sensor = graphql.NewObject(graphql.ObjectConfig{
	Name: "Sensor",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if tempSensor, ok := p.Source.(TempSensor); ok {
					return tempSensor.Name, nil
				}
				return nil, nil
			},
		},
		"rawValue": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if tempSensor, ok := p.Source.(TempSensor); ok {
					return tempSensor.RawValue, nil
				}
				return nil, nil
			},
		},
		"value": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if tempSensor, ok := p.Source.(TempSensor); ok {
					return tempSensor.Value, nil
				}
				return nil, nil
			},
		},
		"target": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if tempSensor, ok := p.Source.(TempSensor); ok {
					return tempSensor.Target, nil
				}
				return nil, nil
			},
		},
		"tolerance": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if tempSensor, ok := p.Source.(TempSensor); ok {
					return tempSensor.Tolerance, nil
				}
				return nil, nil
			},
		},
		"history": &graphql.Field{
			Type: graphql.NewList(SensorHistory),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if tempSensor, ok := p.Source.(TempSensor); ok {
					return tempSensor.History, nil
				}
				return nil, nil
			},
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{

		"sensor": &graphql.Field{
			Type: graphql.NewList(Sensor),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Println(p.Args)
				// TODO: add id filter
				return tempSensors, nil
			},
		},
	},
})
