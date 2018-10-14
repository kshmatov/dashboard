package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type Feed struct {
	event *string
	tstamp time.Time
	pageurl *string
	sourceurl *string
	sourceip *string
	new bool
}

type feeds struct{
	feed Feed
	next *feeds
}

type Feeds struct {
	first *feeds
	last *feeds
	current *feeds
}

func (f *Feeds)Item()*Feed{
	if f.current == nil {
		return nil
	}
	return  &f.current.feed
}

func (f *Feeds)Reset() {
	f.current= f.first
}


func (f *Feeds)First() *Feed{
	f.Reset()
	return f.Item()
}

func (f *Feeds)Next()*Feed{
	if f.current.next == nil {
		return nil
	}
	f.current = f.current.next
	return &f.current.feed
}

func (f *Feeds) add(fd Feed){
	if f.first == nil {
		f.first = &feeds{feed: fd, next: nil}
		f.current = f.first
		f.last = f.first
	} else {
		fs := feeds{feed: fd, next: nil}
		f.current.next = &fs
		f.last = f.current.next
		f.current = f.last
	}
}

func (f *Feeds)addHead(fd Feed) {
	if f.first == nil{
		f.add(fd)
	} else {
		fmt.Printf("Current first: %v, new: %v\n", f.first.feed.sourceurl, fd.sourceurl)
		fs := feeds{feed:fd, next: f.first}
		f.first = &fs
		f.current = &fs
	}
}

func NewFeed(event, pageurl, sourceurl, sourceip string) Feed{
	f := Feed{tstamp:time.Now().UTC().Truncate(time.Millisecond), new:true}
	if len(event) != 0 {
		f.event = &event
	}
	if len(pageurl) != 0 {
		f.pageurl = &pageurl
	}
	if len(sourceurl) != 0 {
		f.sourceurl = &sourceurl
	}
	if len(sourceip) != 0{
		f.sourceip = &sourceip
	}
	return f
}

func (f Feed)Store()error {
	if !f.new {
		return errors.New("Feed can be stored only once")
	}

	sql := "insert into feed(evnt, tstamp, pageurl, sourceurl, sourceip) values($1, $2, $3, $4, $5)"
	_, err := dbConnection.Exec(sql, f.event, f.tstamp, f.pageurl, f.sourceurl, f.sourceip)
	return err
}

func GetFirstN(limit int) (*Feeds, error) {
	if limit < 1 {
		return nil, errors.New("Limit should be greater or equal to 1")
	}
	sql := "select evnt, tstamp, pageurl, sourceurl, sourceip from feed order by tstamp limit $1"
	return fillFeeds(false, sql, limit)
}

func GetFromTS(ts time.Time) (*Feeds, error) {
	sql := "select evnt, tstamp, pageurl, sourceurl, sourceip from feed where tstamp > $1 order by tstamp"
	return fillFeeds(false, sql, ts)
}

func GetLastN(limit int)(*Feeds, error) {
	if limit < 1 {
		return nil, errors.New("Limit should be greater or equal to 1")
	}
	sql := "select evnt, tstamp, pageurl, sourceurl, sourceip from feed order by tstamp desc limit $1"
	return fillFeeds(true, sql, limit)
}

func fillFeeds(revers bool, q string, params ...interface{}) (*Feeds, error) {
	if dbConnection== nil {
		return nil, errors.New("DB connections is not initialized")
	}

	rows, err := dbConnection.Query(q, params...)
	if err != nil {
		return nil, err
	}

	res := Feeds{}
	var evnt, purl, surl, sip *string
	var ts time.Time
	for rows.Next() {
		err := rows.Scan(&evnt, &ts, &purl, &surl, &sip)
		if err != nil {
			return nil, err
		}
		fd := Feed{event:evnt, tstamp:ts, pageurl: purl, sourceip: sip, sourceurl: surl}
		if revers {
			res.addHead(fd)
		} else {
			res.add(fd)
		}
	}
	res.Reset()
	return &res, nil
}

func (f Feed)Event()*string {
	return f.event
}

func (f Feed)PageUrl()*string{
	return f.pageurl
}

func (f Feed)TStamp() time.Time{
	return f.tstamp
}

func (f Feed)SourceIP()*string {
	return f.sourceip
}

func (f Feed)SoureUrl()*string {
	return f.sourceurl
}

func (f Feed)ToJson() []byte{
	m := make(map[string]interface{})
	m["event"] = f.event
	m["tstamp"]	= f.tstamp
	m["pageurl"] = f.pageurl
	m["sourceurl"] = f.sourceurl
	m["sourceip"] = f.sourceip
	res, _:= json.Marshal(m)
	return res
}